<script lang="ts">
  import { Handle, Position, NodeResizer, NodeToolbar, type NodeProps } from '@xyflow/svelte';
  import { JsonView } from '@zerodevx/svelte-json-view';
  import * as Card from "$lib/components/ui/card";
  import type { Writable } from 'svelte/store';
  import * as Pagination from "$lib/components/ui/pagination/index.js";
  import ContextMenu from './ContextMenu.svelte';
  type $$Props = NodeProps;
  let currentPage = 1;
  export let data: { pageCount: Writable<number>, pagesContent: Writable<Array<any>> };
  const { pageCount, pagesContent } = data;
</script>

<ContextMenu id={$$props.id}>
<div class="h-full w-full overflow-hidden svelte-flow__node-default draggable">
  <NodeResizer/>
  <span>{data.label}</span>
  <span>{$$props.id}</span>
  <Handle type="target" position={Position.Left} />
  <Handle type="source" position={Position.Right} />
  <Pagination.Root class="w-full h-[95%]" count={$pageCount} perPage={1} let:pages let:currentPage={currentPage}>
    <Card.Root class="w-full h-full my-2 p-2">
      <Card.Content class="p-0 text-left">
        <JsonView json={$pagesContent[currentPage-1]}/>
      </Card.Content>
    </Card.Root>
    <Pagination.Content class="nodrag">
      <Pagination.Item>
        <Pagination.PrevButton />
      </Pagination.Item>
      {#each pages as page (page.key)}
        {#if page.type === "ellipsis"}
          <Pagination.Item>
            <Pagination.Ellipsis />
          </Pagination.Item>
        {:else}
          <Pagination.Item isVisible={currentPage == page.value}>
            <Pagination.Link {page} isActive={currentPage == page.value}>
              {page.value}
            </Pagination.Link>
          </Pagination.Item>
        {/if}
      {/each}
      <Pagination.Item>
        <Pagination.NextButton />
      </Pagination.Item>
    </Pagination.Content>
  </Pagination.Root>
</div>
</ContextMenu>

