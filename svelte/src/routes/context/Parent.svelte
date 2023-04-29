<script lang="ts">
	import { setContext } from 'svelte';
	import Child from './Child.svelte';
	import Child2 from './Child2.svelte';

	const colorListeners = new Set();
	let colorObj = {
		color: 'red',
    // memory leak.
		listenToColorChange(fn: any) {
			colorListeners.add(fn);
		}
	};

	$: colorListeners.forEach((listener:any) => listener(colorObj.color));
	setContext('color', colorObj);

	let count = 0;
	function onClick() {
		count++;
	}
</script>

<div style="border: 1px solid black;">
	Parent A
	<Child />
	<Child2 color="red" />
</div>

<input type="color" bind:value={colorObj.color} />

<Child2 on:click={onClick} />
{colorObj.color}

Count: {count}
