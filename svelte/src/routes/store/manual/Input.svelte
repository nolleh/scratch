<script lang="ts">
  import { update } from '../../../libs/store';
</script>

<input
  on:input={(e) => {
    update(e.currentTarget.value);
  }}
/>
