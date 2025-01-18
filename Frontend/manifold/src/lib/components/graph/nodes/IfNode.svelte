<script lang="ts">
  import { Handle, Position, type NodeProps } from '@xyflow/svelte';
  import type { Writable } from 'svelte/store';
  import { Input } from '$lib/components/ui/input';
  import ContextMenu from './ContextMenu.svelte';
  type $$Props = NodeProps;
 
  export let data: { 
    path?: Writable<string> 
    pathTwo?: Writable<string> 
    condition?: Writable<string> 
  };
 
  const { path,pathTwo, condition } = data;
  let onInputPathA = (evt:Event) => data.path?.set((evt.target as any)?.value)
  let onInputPathB = (evt:Event) => data.pathTwo?.set((evt.target as any)?.value)
  let onInputCondition = (evt:Event) => data.condition?.set((evt.target as any)?.value)
</script>

<ContextMenu id={$$props.id}>
<div class="h-full w-full svelte-flow__node-default space-y-2 draggable">
  <span>{data.label}</span>
  <span>{$$props.id}</span>
  <Handle id="recordA" class="top-1/4" type="target" position={Position.Left} />
  <Handle id="recirdB" class="bottom-1/4" style="top: unset;" type="target" position={Position.Left} />
  <br/>
  <div>
    <span> Path record A</span>
  </div>
  <Input
    id="recordAPath"
    class="nodrag h-full w-full resize-none"
    on:input={onInputPathA}
    value={$path}
  />
  <div>
    <span>Path record B</span>
  </div>
  <Input
    id="recordBPath"
    class="nodrag h-full w-full resize-none"
    on:input={onInputPathB}
    value={$pathTwo}
  />
  <div>
    <span>Condition</span>
  </div>
  <Input
    id="Condition"
    class="nodrag h-full w-full resize-none"
    on:input={onInputCondition}
    value={$condition}
  />
  <Handle class="top-1/4" id="ifTrue" type="source" position={Position.Right} />
  <Handle class="bottom-1/4" style="top: unset;" id="else" type="source" position={Position.Right} />
</div>
</ContextMenu>
