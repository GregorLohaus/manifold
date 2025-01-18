<script lang="ts">
  import RegisterForm from "./RegisterForm.svelte"
  export let data
</script>
<h1 class="text-2xl mb-3">Register</h1>
<RegisterForm data={data.form}/>
