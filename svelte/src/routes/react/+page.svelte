<script lang="ts">
  import Child from './Child.svelte';

  let componentTitle = 'child';

  $: data = {
    title: componentTitle
  };

  function click() {
    console.log('clicked');
    componentTitle = 'changed prop';
  }
</script>

<button on:click={click}> click 2 change </button>
<Child title= "with wrap obj" prop={data.title} />
<Child title= "without wrap obj" prop={componentTitle} />
