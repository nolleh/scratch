
<script lang='ts'>
  import { onMount } from 'svelte';
  import { valueStore } from '../../../libs/store2';
  let _value = '';
  onMount(() => {
    return valueStore.subscribe(value => {
      _value = value;
    })
  });
</script>

<h1>
	value: {_value}
</h1>
