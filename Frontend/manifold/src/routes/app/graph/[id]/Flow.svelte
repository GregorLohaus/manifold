<script lang="ts">
	import { writable, type Writable } from 'svelte/store';
	import * as nt from '$lib/constants/nodetypes';
	import { getRandomHex } from '$lib/get_random_hex';
	import {
		SvelteFlow,
		type Node,
		type Edge,
		Controls,
		MiniMap,
		Background,
		useSvelteFlow
	} from '@xyflow/svelte';
	import TextInputNode from '$lib/components/graph/nodes/TextInputNode.svelte';
	import ParseCsvNode from '$lib/components/graph/nodes/ParseCsvNode.svelte';
	import AddNode from '$lib/components/graph/nodes/AddNode.svelte';
	import Toolbar from './Toolbar.svelte';
	import RemoveNode from '$lib/components/graph/nodes/RemoveNode.svelte';
	import MergeNode from '$lib/components/graph/nodes/MergeNode.svelte';
	import SetNode from '$lib/components/graph/nodes/SetNode.svelte';
	import SelectNode from '$lib/components/graph/nodes/SelectNode.svelte';
	import IfNode from '$lib/components/graph/nodes/IfNode.svelte';
	import DataViewNode from '$lib/components/graph/nodes/DataViewNode.svelte';
	import './flow.css';
	import { getContext } from 'svelte';
	import { DefaultApiPath } from '$lib/default_api_url';
	let graphid: Writable<string> = getContext('graphid');
	const nodes = writable<Node[]>([]);
	const edges = writable<Edge[]>([]);
	$: {
		console.log($graphid);
		if ($graphid !== 'new') {
			fetch(DefaultApiPath(`graph/${$graphid}`), {
				method: 'GET',
				credentials: 'include',
				headers: {
					'Content-Type': 'application/json'
				}
			}).then((res) => {
				res.text().then((t) => {
					if (res.ok) {
						let graph = JSON.parse(t);
						graph.nodes?.map((n:any) => {
							console.log(n)
							for (let [k,v] of Object.entries(n.data)) {
								console.log(v)
								if (k !== "label") {
									n.data[k] = writable(v)
								}
							}
						})
						nodes.set(graph.nodes);
						edges.set(graph.edges);
					}
				});
			});
		}
	}
	const { screenToFlowPosition } = useSvelteFlow();
	const nodeTypes = {
		[nt.TEXT_INPUT]: TextInputNode,
		[nt.ADD]: AddNode,
		[nt.REMOVE]: RemoveNode,
		[nt.MERGE]: MergeNode,
		[nt.SET]: SetNode,
		[nt.SELECT]: SelectNode,
		[nt.IF]: IfNode,
		[nt.DATAVIEW]: DataViewNode,
		[nt.PARSE_CSV]: ParseCsvNode
	};
	const onDragOver = (event: DragEvent) => {
		event.preventDefault();

		if (event.dataTransfer) {
			event.dataTransfer.dropEffect = 'move';
		}
	};
	const onDrop = (event: DragEvent) => {
		event.preventDefault();

		if (!event.dataTransfer) {
			return null;
		}
		const type = event.dataTransfer.getData('nodeType');
		const position = screenToFlowPosition({
			x: event.clientX,
			y: event.clientY
		});
		let newNode = {
			id: `${getRandomHex(6)}`,
			type,
			position,
			data: { label: `${type}` },
			origin: [0.5, 0.0]
		} satisfies Node;
		switch (type) {
			case nt.TEXT_INPUT:
				(newNode.data as any).text = writable('');
				break;
			case nt.PARSE_CSV:
				break;
			case nt.ADD:
				(newNode.data as any).path = writable('');
				break;
			case nt.REMOVE:
				(newNode.data as any).path = writable('');
				break;
			case nt.MERGE:
				break;
			case nt.SET:
				(newNode.data as any).path = writable('');
				break;
			case nt.SELECT:
				(newNode.data as any).path = writable('');
				break;
			case nt.IF:
				(newNode.data as any).path = writable('');
				(newNode.data as any).pathTwo = writable('');
				(newNode.data as any).condition = writable('');
				break;
			case nt.DATAVIEW:
				(newNode.data as any).pageCount = writable(1);
				(newNode.data as any).pagesContent = writable([{ yourData: 'here' }]);
				break;
		}
		$nodes.push(newNode);
		$nodes = $nodes;
	};
</script>

<Toolbar />
<SvelteFlow {nodeTypes} {nodes} {edges} on:dragover={onDragOver} on:drop={onDrop}>
	<Background
		bgColor="hsl(var(--background) / var(--tw-bg-opacity))"
		patternColor="hsl(var(--foreground) / var(--tw-text-opacity))"
	></Background>
	<Controls />
	<MiniMap class="rounded-md" pannable zoomable position="top-right" height={120} />
</SvelteFlow>
