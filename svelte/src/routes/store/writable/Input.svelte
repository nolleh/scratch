<script lang="ts">
  import { valueStore } from '../../../libs/store2';
</script>

<input
  on:input={(e) => {
    valueStore.set(e.currentTarget.value);
  }}
/>
