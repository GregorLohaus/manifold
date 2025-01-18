<script lang="ts">
  import { Handle, Position, NodeResizer, type NodeProps } from '@xyflow/svelte';
  import type { Writable } from 'svelte/store';
  import { Input } from '$lib/components/ui/input';
  import ContextMenu from './ContextMenu.svelte'; 
  type $$Props = NodeProps;
 
  export let data: { path?: Writable<string> };
 
  const { path } = data;
  let onInput = (evt:Event) => data.path?.set((evt.target as any)?.value)
</script>

<ContextMenu id={$$props.id}>
<div class="h-full w-full svelte-flow__node-default draggable">
  <span>{data.label}</span>
  <span>{$$props.id}</span>
  <Handle id="record" class="top-1/4" type="target" position={Position.Left} />
  <Handle style="top:unset;" class="bottom-1/4" id="value" type="target" position={Position.Left} />
  <Input
    class="nodrag h-full w-full resize-none"
    on:input={onInput}
    value={$path}
  />
  <Handle type="source" position={Position.Right} />
</div>
</ContextMenu>
