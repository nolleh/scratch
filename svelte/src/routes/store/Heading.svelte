
<script lang='ts'>
  import { onMount } from 'svelte';
  import { value, subscribe } from '../../libs/store';
  let _value = '';
  onMount(() => {
    return subscribe(() => {
      _value = value;
    });
  });
</script>

<h1>
	value: {_value}
</h1>
