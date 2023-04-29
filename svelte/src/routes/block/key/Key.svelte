<script>
	import { fade } from 'svelte/transition';
	let count = 0;
</script>

<!-- recreate component -->
<!-- expression change, intput gone -->
{#key count}
	<input />
	<div in:fade>{count}</div>
{/key}

<button on:click={() => count++}> {count} </button>
