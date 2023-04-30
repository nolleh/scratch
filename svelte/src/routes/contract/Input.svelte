<script>
  import { valueStore } from '../../libs/store2';
</script>

<!-- <input -->
<!--   on:input={(e) => { -->
<!--     $valueStore = e.currentTarget.value; -->
<!--   }} -->
<!-- /> -->

<input bind:value={$valueStore}/>
