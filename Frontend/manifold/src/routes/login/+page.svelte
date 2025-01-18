<script lang="ts">
  import LoginForm from './LoginForm.svelte';
  export let data
</script>

<h1 class="text-2xl mb-3">Login</h1>
<LoginForm data={data.form}/>
