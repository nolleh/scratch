<script lang="ts">
  export let title: string;
  export let prop: string;

  $: bad = () => prop
  $: good = (t2: string) => {
    return t2;
  };
</script>

<h3>{title}</h3>
{bad()} <br/>
{good(prop)}
