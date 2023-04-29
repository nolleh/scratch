<script lang="ts">
	import { getContext } from 'svelte';
	import GrandChildA from './GrandChildA.svelte';
	import GrandChildB from './GrandChildB.svelte';

  const color:string = getContext('color');

</script>

<GrandChildA {color}/>
<GrandChildB/>
