<script lang="ts">
  import { Handle, Position, NodeResizer, type NodeProps } from '@xyflow/svelte';
  import type { Writable } from 'svelte/store';
  import { Textarea } from "$lib/components/ui/textarea/index.js";
  import ContextMenu from './ContextMenu.svelte';
  type $$Props = NodeProps;
 
  export let data: { text?: Writable<string> };
 
  const { text } = data;
  let onInput = (evt:Event) => data.text?.set((evt.target as any)?.value)
</script>

<ContextMenu id={$$props.id}>
<div class="h-full overflow-hidden w-full svelte-flow__node-default draggable">
  <NodeResizer/>
  <span>{data.label}</span>
  <span>{$$props.id}</span>
  <Textarea
    class="nodrag h-[95%] w-full resize-none"
    on:input={onInput}
    value={$text}
  />
  <Handle type="source" position={Position.Right} />
</div>
</ContextMenu>
