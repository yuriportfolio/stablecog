<script lang="ts">
	import ScBar from '$components/ScBar.svelte';
	import LL from '$i18n/i18n-svelte';
	import LinkButton from '$components/buttons/NoBgButton.svelte';
	import IconBack from '$components/icons/IconBack.svelte';
	import { isTouchscreen } from '$ts/stores/isTouchscreen';
	import { windowHeight } from '$ts/stores/window';
	import { navbarHeight } from '$ts/stores/navbarHeight';

	export let toc: string;
	export { classes as class };
	let classes = 'hidden lg:flex';
</script>

<div
	style={$windowHeight && $navbarHeight
		? `max-height: calc(${$windowHeight - $navbarHeight}px); top: calc(${$navbarHeight}px);`
		: 'max-height: calc(100vh-5rem); max-height: calc(100svh-5rem);'}
	class="w-72 self-stretch flex flex-col sticky top-24 mt-1 overflow-hidden {classes}"
>
	<div class="h-full absolute right-0 top-0 py-4">
		<div class="h-full w-2px bg-c-on-bg/5 rounded-full" />
	</div>
	<div
		class="w-full h-full items-start justify-start flex flex-col 
		self-stretch"
	>
		<div class="w-full flex justify-center mb-1 px-5 mt-5">
			<LinkButton href="/blog" prefetch={true}>
				<div class="flex items-center justify-center gap-2.5 px-2 py-1">
					<IconBack
						class="w-6 h-6 transform transition text-c-on-bg/50 group-hover:-translate-x-1
						{!$isTouchscreen ? 'group-hover:text-c-primary' : ''}"
					/>
					<p class="font-bold">{$LL.Blog.BackToBlogButton()}</p>
				</div>
			</LinkButton>
		</div>
		<div class="w-full px-5">
			<ScBar />
		</div>
		<div class="w-full px-5 mt-5">
			<div class="w-full h-2px bg-c-on-bg/5 rounded-full px-5" />
		</div>
		<nav class="w-full toc max-h-full overflow-auto pt-5 pb-16 px-5">
			{@html toc}
		</nav>
	</div>
</div>
