<script lang="ts">
	import GrandChildB from './GrandChildB.svelte';
	export let color: string = 'red';

	let show = true;
</script>

child2: >> {color}
<label>
	<input type="checkbox" bind:checked={show} />
	Show
</label>

{#if show}
	<!-- event forwarding -->
	<GrandChildB on:click />
{/if}
