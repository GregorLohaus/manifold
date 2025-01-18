<script lang="ts">
  import VerifyForm from './VerifyForm.svelte';
  export let data
</script>

<h1 class="text-2xl mb-3">Verify</h1>
<VerifyForm data={data.form}/>
