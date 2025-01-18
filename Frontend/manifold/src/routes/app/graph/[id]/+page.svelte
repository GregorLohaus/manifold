<script lang="ts">
  import { SvelteFlowProvider } from '@xyflow/svelte';

  import Flow from './Flow.svelte';
	import { writable, type Writable } from 'svelte/store';
	import { getContext } from 'svelte';
  type pagedata = { id : string }
  export let data:pagedata;
	let graphid: Writable<string> = getContext('graphid');
  graphid.set(data.id)
</script>

<SvelteFlowProvider>
  <Flow />
</SvelteFlowProvider>

