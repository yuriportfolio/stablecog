<script lang="ts">
	import { isTouchscreen } from '$ts/stores/isTouchscreen';

	export let hoverFrom: 'top' | 'left' | 'bottom' | 'right' = 'bottom';
	export let size: 'lg' | 'md' = 'lg';
	export let color: 'primary' | 'on-bg' | 'bg-secondary' = 'primary';
	export let hovered = false;
	export let noPadding = false;
</script>

<div class="w-full h-full absolute left-0 top-0 {!noPadding ? 'p-0.5' : ''}">
	<div
		class="w-full h-full relative z-0 overflow-hidden {size === 'md' ? 'rounded-lg' : 'rounded-xl'}"
	>
		<div
			class="w-full h-full origin-left transition transform {hovered
				? 'translate-x-0 translate-y-0'
				: hoverFrom === 'top'
				? '-translate-y-full'
				: hoverFrom === 'right'
				? 'translate-x-1/2 opacity-0'
				: hoverFrom === 'left'
				? '-translate-x-1/2 opacity-0'
				: 'translate-y-full'}
      {color === 'on-bg'
				? 'bg-c-on-bg/10'
				: color === 'bg-secondary'
				? 'bg-c-bg-secondary'
				: 'bg-c-primary/15'} {size === 'md' ? 'rounded-lg' : 'rounded-xl'} {!$isTouchscreen
				? 'group-hover:translate-x-0 group-hover:translate-y-0 group-hover:opacity-100'
				: ''}"
		/>
	</div>
</div>
