<script lang="ts">
	import { goto } from "$app/navigation";
  import * as Menubar from "$lib/components/ui/menubar";
	import { DefaultApiPath } from "$lib/default_api_url";
	import { redirect } from "@sveltejs/kit";
  import { useNodes,useEdges,type Node, type Edge } from "@xyflow/svelte";
	import { getContext } from "svelte";
	import { type Writable } from "svelte/store";
  const nodes = useNodes();
  const edges = useEdges();
  type Graph = { nodes: Array<Node> ,edges: Array<Edge>}
  const onSave = async () => {
    let graphid: Writable<string> = getContext('graphid');
    let graph:Graph = { nodes: [] ,edges: []};
    let us = nodes.subscribe((v) => { 
      graph.nodes = v.map((vv) => {
        let { id,origin,position,selected,type } = vv 
          return {id: id,origin:origin,position:position,selected: selected,type:type,data:{}}
        }
      ) 
    });
    us();
    us = edges.subscribe((v) => { graph.edges = v });
    us();
    $nodes.forEach((node) => {
      let gN:Node|undefined = undefined
      graph.nodes.forEach((n) => {
        if (n.id == node.id) {
          gN = n 
        }
      })
      if (typeof gN === 'undefined') {
        return
      }
      for (let [k,v] of Object.entries(node.data)) {
        if (typeof (v as any).subscribe === 'function') {
          let vT = (v as Writable<any>)
          let unsubscribe = vT.subscribe((e:any) => {
            (gN as Node).data[k] = structuredClone(e)
          }) 
          unsubscribe()
        }
        else {
          (gN as Node).data[k] = v
        }
      }
    })
    let id = "new"
    let sub = graphid.subscribe((v)=>{id = v})
    sub()
    let res = await fetch(DefaultApiPath(`graph/${id}`),{
      method: "POST",
      body: JSON.stringify(graph),
      credentials: "include",
      headers: {
        "Content-Type" : "application/json"
      }
    })
    let body = await res.text()
    if (JSON.parse(body)?.id != null) {
      goto(`/app/graph/${JSON.parse(body)?.id}`)
    }
  }
</script>

<Menubar.Item on:click={onSave}>
    Save
</Menubar.Item>

