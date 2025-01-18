<script lang="ts">
	import * as Collapsible from '$lib/components/ui/collapsible';
	import Button from '$lib/components/ui/button/button.svelte';
	import { CaretSortSolid, CodeBranchSolid } from 'flowbite-svelte-icons';
	import { writable, type Writable } from 'svelte/store';
	import { DefaultApiPath } from '$lib/default_api_url';
	import { goto, invalidate, invalidateAll } from '$app/navigation';
	import { getContext } from 'svelte';
	let graphids = writable([]);
	fetch(DefaultApiPath('graphs'), {
		method: 'GET',
		credentials: 'include',
		headers: {
			'Content-Type': 'application/json'
		}
	}).then((res) => {
		res.text().then((t) => {
			if (res.ok) {
				console.log(t);
				graphids.set(JSON.parse(t));
			}
		});
	});
	let graphid: Writable<string> = getContext('graphid');
	const reload= (id:any) => {
		graphid.set(id.id)
		goto(`/app/graph/${id.id}`)
	}
</script>

<Collapsible.Root>
	<div class="flex items-center justify-between w-full space-y-2">
		<h4 class="w-full flex flex-row justify-between items-center">
			<CodeBranchSolid class="mr-2" />
			Graphs
			<Collapsible.Trigger asChild let:builder>
				<Button builders={[builder]} variant="ghost" size="sm" class="ml-3">
					<CaretSortSolid />
				</Button>
			</Collapsible.Trigger>
		</h4>
	</div>
	<Collapsible.Content class="space-y-2">
		<div class="h-2"></div>
		<a
			href="/app/graph/new"
			class="bg-primary text-secondary flex items-center justify-center px-2 py-1 rounded-md"
		>
			New
		</a>
		{#each $graphids as id}
			<a
				on:click={() => reload(id)}
				class="bg-primary text-secondary flex items-center justify-center px-2 py-1 rounded-md"
			>
				{id.id}
			</a>
		{/each}
	</Collapsible.Content>
</Collapsible.Root>
