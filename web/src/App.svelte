<div class="container">
	{#await request}
		<h1>LOADING</h1>
	{:then json}
		{#each json as row}
			{#each row as block}
				<span contenteditable="false" bind:innerHTML="{block_list[block]}" />
			{/each}<br>
		{/each}
	{:catch err}
		<h1>{err}</h1>
	{/await}
</div>

<script>

let block_list = ["&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;","◯","⚫","██"]

async function apiCall() {
	let uuid = "99ae4996-c5a4-4b14-a78b-813ccdc358cd"
	const res = await fetch(`http://localhost:3000/getlayout/${uuid}`);
	const text = await res.json();

	if (res.ok) {
		return text;
	} else {
		throw new Error(text);
	}
}

let request = apiCall()

</script>

<style>
</style>