<script lang="ts">
  import * as Menubar from "$lib/components/ui/menubar";
  import { Panel,useNodes,useEdges,type Node, type Edge } from "@xyflow/svelte";
  import * as nt from "$lib/constants/nodetypes"
  import ToolbarSave from "./ToolbarSave.svelte";
  const onDragStart = (event:any, type: string) => {
    if (!event.dataTransfer) {
      return null;
    }
    event.dataTransfer.setData('nodeType', type);
    event.dataTransfer.effectAllowed = 'move';
  };

</script>

<Panel position="top-center">
<Menubar.Root>
  <Menubar.Menu>
    <Menubar.Trigger>
      File
    </Menubar.Trigger>
    <Menubar.Content>
      <ToolbarSave/>
      <Menubar.Item>
          New
      </Menubar.Item>
    </Menubar.Content>
  </Menubar.Menu>
  <Menubar.Menu>
    <Menubar.Trigger>
      Input Nodes
    </Menubar.Trigger>
    <Menubar.Content>
      <Menubar.Item>
        <div 
          on:dragstart={(event) => onDragStart(event, nt.TEXT_INPUT)}
          draggable={true}
          role="definition"
        >
          Textinput
        </div>
      </Menubar.Item>
    </Menubar.Content>
  </Menubar.Menu>
  <Menubar.Menu>
    <Menubar.Trigger>
      Parser Nodes
    </Menubar.Trigger>
    <Menubar.Content>
      <Menubar.Item>
        <div 
          on:dragstart={(event) => onDragStart(event, nt.PARSE_CSV)}
          draggable={true}
          role="definition"
        >
          Parse CSV
        </div>
      </Menubar.Item>
    </Menubar.Content>
  </Menubar.Menu>
  <Menubar.Menu>
    <Menubar.Trigger>
      Transform Nodes
    </Menubar.Trigger>
    <Menubar.Content>
      <Menubar.Item>
        <div 
          on:dragstart={(event) => onDragStart(event, nt.ADD)}
          draggable={true}
          role="definition"
        >
          Add
        </div>
      </Menubar.Item>
      <Menubar.Item>
        <div 
          on:dragstart={(event) => onDragStart(event, nt.REMOVE)}
          draggable={true}
          role="definition"
        >
          Remove
        </div>
      </Menubar.Item>
      <Menubar.Item>
        <div 
          on:dragstart={(event) => onDragStart(event, nt.MERGE)}
          draggable={true}
          role="definition"
        >
          Merge
        </div>
      </Menubar.Item>
      <Menubar.Item>
        <div 
          on:dragstart={(event) => onDragStart(event, nt.SET)}
          draggable={true}
          role="definition"
        >
          Set
        </div>
      </Menubar.Item>
      <Menubar.Item>
        <div 
          on:dragstart={(event) => onDragStart(event, nt.SELECT)}
          draggable={true}
          role="definition"
        >
          Select
        </div>
      </Menubar.Item>
    </Menubar.Content>
  </Menubar.Menu>
  <Menubar.Menu>
    <Menubar.Trigger>
      Logic Nodes
    </Menubar.Trigger>
    <Menubar.Content>
      <Menubar.Item>
        <div 
          on:dragstart={(event) => onDragStart(event, nt.IF)}
          draggable={true}
          role="definition"
        >
          If
        </div>
      </Menubar.Item>
    </Menubar.Content>
  </Menubar.Menu>
  <Menubar.Menu>
    <Menubar.Trigger>
      Output Nodes
    </Menubar.Trigger>
    <Menubar.Content>
      <Menubar.Item>
        <div 
          on:dragstart={(event) => onDragStart(event, nt.DATAVIEW)}
          draggable={true}
          role="definition"
        >
          Dataview
        </div>
      </Menubar.Item>
    </Menubar.Content>
  </Menubar.Menu>
</Menubar.Root>
</Panel>

