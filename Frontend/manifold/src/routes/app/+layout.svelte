<script lang="ts">
	import * as Resizable from '$lib/components/ui/resizable';
	import Menu from './Menu.svelte';
</script>

<Resizable.PaneGroup direction="horizontal">
	<Resizable.Pane
		class="flex flex-col flex-wrap justify-center content-center h-screen"
		minSize={0}
		defaultSize={15}
	>
		<Menu />
	</Resizable.Pane>
	<Resizable.Handle withHandle/>
	<Resizable.Pane class="flex flex-col flex-wrap justify-center content-center h-screen p-5">
		<slot />
	</Resizable.Pane>
</Resizable.PaneGroup>
