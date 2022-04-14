<input bind:this={inputEl} type="text">
<button on:click={handleClick}>GET</button>

{#if request != undefined}
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
			<h4>{err}</h4>
		{/await}
	</div>
{/if}
<script>

export let request;
let inputEl;
let block_list = ["&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;","◯","⚫","██"]

async function apiCall() {
	let uuid = inputEl.value;
	const res = await fetch(`http://localhost:3000/getlayout/${uuid}`);
	const text = await res.json();

	if (res.ok) {
		return text;
	} else {
		throw new Error(text);
	}
}

const handleClick = () => {
	request = apiCall()
}

</script>

<style>
</style>