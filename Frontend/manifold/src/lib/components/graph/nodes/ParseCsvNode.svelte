<script lang="ts">
  import { Handle, Position, type NodeProps } from '@xyflow/svelte';
  import ContextMenu from './ContextMenu.svelte';
  type $$Props = NodeProps;
  export let data:any;
</script>

<ContextMenu id={$$props.id}>
<div class="svelte-flow__node-default draggable">
  <span>{data.label}</span>
  <span>{$$props.id}</span>
  <Handle type="target" position={Position.Left} />
  <Handle type="source" position={Position.Right} />
</div>
</ContextMenu>
