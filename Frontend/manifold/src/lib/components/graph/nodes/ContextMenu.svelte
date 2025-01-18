<script lang="ts">
  import { useEdges, useNodes } from "@xyflow/svelte";
  import * as ContextMenu from "$lib/components/ui/context-menu/index.js";
  export let id:string;
  const nodes = useNodes()
  const edges = useEdges()
  const onDelete = async() => {
    $nodes = $nodes.filter((node) => node.id !== id);
    $edges = $edges.filter((edge) => edge.source !== id && edge.target !== id);
  }
</script>

<ContextMenu.Root >
  <ContextMenu.Trigger class="h-full w-full">
    <slot/>
  </ContextMenu.Trigger>
  <ContextMenu.Content>
    <ContextMenu.Item on:click={onDelete}>
      Delete
    </ContextMenu.Item>
  </ContextMenu.Content>
</ContextMenu.Root>

