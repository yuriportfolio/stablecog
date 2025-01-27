<script lang="ts">
	import { page } from '$app/stores';
	import LL from '$i18n/i18n-svelte';
	import InsufficientCreditsBadge from '$components/badges/InsufficientCreditsBadge.svelte';
	import Button from '$components/buttons/Button.svelte';
	import ClearButton from '$components/buttons/ClearButton.svelte';
	import { autoresize } from '$ts/actions/textarea/autoresize';
	import { maxPromptLength, maxSeed } from '$ts/constants/main';
	import { formatPrompt } from '$ts/helpers/formatPrompt';
	import { generateSSEId } from '$ts/helpers/generateSSEId';
	import {
		generationGuidanceScale,
		generationHeight,
		generationInferenceSteps,
		generationInitImageFiles,
		generationInitImageFilesState,
		generationInitImageStrength,
		generationInitImageUrl,
		generationModelId,
		generationNumOutputs,
		generationSchedulerId,
		generationSeed,
		generationWidth,
		generationNegativePrompt,
		prompt,
		generationPrompt
	} from '$ts/stores/generationSettings';
	import { isTouchscreen } from '$ts/stores/isTouchscreen';
	import {
		generations,
		queueInitialGenerationRequest,
		type TInitialGenerationRequest
	} from '$ts/stores/user/generation';
	import { sseId } from '$ts/stores/user/sse';
	import { userSummary } from '$ts/stores/user/summary';
	import { maxOngoingGenerationOutputsCount } from '$routes/admin/create/constants';
	import IconGenerationSettings from '$components/icons/IconGenerationSettings.svelte';
	import IconButton from '$components/buttons/IconButton.svelte';
	import { windowWidth } from '$ts/stores/window';
	import { mdBreakpoint } from '$components/generationFullScreen/constants';
	import type { TCreatePageData } from '$routes/admin/create/+page.server';
	import { onMount } from 'svelte';
	import { isValue } from '$ts/helpers/isValue';
	import { isSuperAdmin } from '$ts/helpers/admin/roles';
	import type { TIsReadyMap } from '$routes/admin/create/types';

	export let openSignInModal: () => void;
	export let openSettingsPanelModal: (() => void) | undefined = undefined;
	export let serverData: TCreatePageData;
	export let isReadyMap: TIsReadyMap;

	let promptInputElement: HTMLTextAreaElement;
	let isCheckCompleted = false;

	generationPrompt.set(serverData.prompt !== null ? serverData.prompt : undefined);

	$: [$generationPrompt], withCheck(setLocalPrompt);

	$: doesntHaveEnoughCredits =
		isCheckCompleted &&
		$userSummary &&
		$userSummary.total_remaining_credits < Number($generationNumOutputs);

	$: showClearPromptInputButton = $generationPrompt !== undefined && $generationPrompt !== '';
	$: promptInputPlaceholder = $LL.Home.PromptInput.Placeholder();

	$: ongoingGenerationOutputsCount = $generations
		.map((g) => g.status)
		.filter((s) => s !== 'succeeded' && s !== 'failed').length;
	$: maxOngoingGenerationOutputsCountReached = isSuperAdmin($userSummary?.roles)
		? false
		: ongoingGenerationOutputsCount >= $maxOngoingGenerationOutputsCount;

	async function onPromptFormSubmitted() {
		if (maxOngoingGenerationOutputsCountReached) {
			return;
		}
		if (!$page.data.session?.user.id) {
			openSignInModal();
			return;
		}
		if ($generationInitImageFilesState === 'uploading') {
			return;
		}
		if ($generationPrompt) {
			generationPrompt.set(formatPrompt($generationPrompt));
		}
		if ($generationNegativePrompt) {
			generationNegativePrompt.set(formatPrompt($generationNegativePrompt));
		}
		if (!$generationPrompt) {
			generationPrompt.set(promptInputPlaceholder);
		}
		queueGeneration();
		if ($windowWidth < mdBreakpoint) {
			promptInputElement.blur();
		}
	}

	async function queueGeneration() {
		if (!$generationPrompt) {
			console.log("No prompt, can't create generation");
			return;
		}
		if (!$sseId) {
			console.log("No SSE ID, can't create generation");
			return;
		}
		const initialRequestProps: TInitialGenerationRequest = {
			prompt: {
				id: 'prompt',
				text: $generationPrompt
			},
			negative_prompt: $generationNegativePrompt
				? {
						id: 'negative_prompt',
						text: $generationNegativePrompt
				  }
				: undefined,
			model_id: $generationModelId,
			scheduler_id: $generationSchedulerId,
			width: Number($generationWidth),
			height: Number($generationHeight),
			init_image_url: $generationInitImageUrl,
			init_image_file: $generationInitImageFiles,
			prompt_strength:
				$generationInitImageUrl && $generationInitImageStrength !== undefined
					? 1 - Number($generationInitImageStrength)
					: undefined,
			guidance_scale: Number($generationGuidanceScale),
			inference_steps: Number($generationInferenceSteps),
			seed:
				$generationSeed !== undefined && $generationSeed !== null && $generationSeed !== ''
					? Number($generationSeed)
					: Math.round(Math.random() * maxSeed),
			num_outputs: Number($generationNumOutputs),
			output_image_extension: 'jpeg',
			process_type: 'generate',
			stream_id: $sseId,
			ui_id: generateSSEId(),
			submit_to_gallery: false
		};
		queueInitialGenerationRequest(initialRequestProps);
	}

	function clearPrompt() {
		generationPrompt.set('');
		promptInputElement.value = '';
		promptInputElement.blur();
		promptInputElement.focus();
	}

	function setLocalPrompt() {
		prompt.set(
			$generationPrompt !== '' && $generationPrompt !== undefined ? $generationPrompt : ''
		);
	}

	function withCheck(fn: () => void) {
		if (!isCheckCompleted) return;
		fn();
	}

	onMount(() => {
		isCheckCompleted = false;
		if (!isValue(serverData.prompt) && isValue($prompt) && $prompt !== null) {
			generationPrompt.set($prompt);
		}
		isCheckCompleted = true;
		isReadyMap.promptBar = true;
	});
</script>

<div class="w-full flex justify-center pb-3 pl-2 md:p-0">
	<form
		on:submit|preventDefault={onPromptFormSubmitted}
		class="w-full max-w-7xl flex flex-row md:gap-3 items-center"
	>
		<div class="flex-1 flex relative group">
			<textarea
				use:autoresize={{ maxRows: 3, placeholder: promptInputPlaceholder }}
				bind:this={promptInputElement}
				bind:value={$generationPrompt}
				on:keypress={(e) => {
					if (e.key === 'Enter') {
						e.preventDefault();
						onPromptFormSubmitted();
					}
				}}
				on:input={() => {
					if (
						$generationPrompt !== undefined &&
						$generationPrompt !== null &&
						$generationPrompt.length > maxPromptLength
					) {
						generationPrompt.set($generationPrompt.slice(0, maxPromptLength));
					}
				}}
				disabled={!isCheckCompleted}
				placeholder={promptInputPlaceholder}
				enterkeyhint="go"
				rows="1"
				style="transition: height 0.1s cubic-bezier(0.4, 0, 0.2, 1), box-shadow 0.2s cubic-bezier(0.4, 0, 0.2, 1), padding 0.2s cubic-bezier(0.4, 0, 0.2, 1)"
				class="w-full text-base bg-c-bg-secondary shadow-lg pr-12 md:pr-17 hide-scrollbar shadow-c-shadow/[var(--o-shadow-normal)] 
						scroll-smooth resize-none transition relative pl-3 md:pl-6 py-3.5 md:py-5 rounded-xl 
						focus:ring-2 focus:ring-c-primary/30 ring-0 ring-c-primary/20 placeholder:text-c-on-bg/40 {!$isTouchscreen
					? 'enabled:hover:ring-2'
					: ''} text-c-on-bg {!$isTouchscreen ? 'group-hover:ring-2' : ''}"
			/>
			<ClearButton show={showClearPromptInputButton} onClick={clearPrompt} />
		</div>
		{#if openSettingsPanelModal}
			<IconButton name="Generation Settings" onClick={openSettingsPanelModal}>
				<div class="p-1">
					<IconGenerationSettings
						class="{!$isTouchscreen
							? 'group-hover/iconbutton:text-c-primary'
							: ''} text-c-on-bg transition"
					/>
				</div>
			</IconButton>
		{/if}
		<div class="w-full md:w-auto md:min-w-[9.5rem] relative hidden md:block">
			<Button
				disabled={!isCheckCompleted ||
					(doesntHaveEnoughCredits && $page.data.session?.user.id !== undefined)}
				uploading={$generationInitImageFilesState === 'uploading'}
				loading={maxOngoingGenerationOutputsCountReached}
				withSpinner
				fadeOnDisabled={isCheckCompleted}
				class="w-full flex flex-col relative"
			>
				<p>
					{$LL.Home.GenerateButton()}
				</p>
			</Button>
			{#if doesntHaveEnoughCredits && $userSummary && $page.data.session?.user.id}
				<InsufficientCreditsBadge
					neededCredits={Number($generationNumOutputs)}
					remainingCredits={$userSummary.total_remaining_credits}
				/>
			{/if}
		</div>
	</form>
</div>
