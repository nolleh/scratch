<script lang="ts">
	// export let color: string = 'red';
  import { getContext } from 'svelte';

  let { color, listenToColorChange }: any = getContext('color');
  listenToColorChange((newColor:string) => {
    color = newColor;
  });
</script>

<div style="color: {color}">
  GrandChild B: {color}
  <button on:click>Click Me</button>
</div>
