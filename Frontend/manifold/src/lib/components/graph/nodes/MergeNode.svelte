<script lang="ts">
  import { Handle, Position, NodeToolbar, useSvelteFlow, type NodeProps } from '@xyflow/svelte';
	import ContextMenu from './ContextMenu.svelte';
  type $$Props = NodeProps;
  export let data:any;
</script>

<ContextMenu id={$$props.id}>
<div class="svelte-flow__node-default draggable">
  <span>{data.label}</span>
  <span>{$$props.id}</span>
  <Handle class="top-1/4" id="recordA" type="target" position={Position.Left} />
  <Handle style="top:unset;" id="recordB" type="target" position={Position.Left} />
  <Handle type="source" position={Position.Right} />
</div>
</ContextMenu>
