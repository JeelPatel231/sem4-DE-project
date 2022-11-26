<div class="flex-container">
	<input class="search" bind:this={inputEl} value="90c88825-669b-4bf7-9320-20771609fb0e" type="text">
	<div class="getButton" on:click={handleClick}>GET</div>
</div>
	
{#if request != undefined}
	<div class="container">
		{#await request}
			<h1>LOADING</h1>
		{:then json}
			{#each json as row,row_index}
				<div class="row">
					{#each row as block,col_index}
					<span >
						<svelte:component 
							on:seatDataChange={(checked)=>{changeData(row_index,col_index,checked)}}
							this={block_list[block]}
							checked={block == 2}
						/>
					</span>
					{/each}
				</div>
			{/each}
		{:catch err}
			<h4>{err}</h4>
		{/await}
	</div>
{/if}

<!-- <button on:click={setData}>SET DATA</button> -->

<script>
import Seat from "./components/seat.svelte"
import Table from "./components/table.svelte"
import Space from "./components/space.svelte";
import { onMount } from "svelte";

let clientData;
let request;
let inputEl;
let block_list = [Space,Seat,Seat,Table]

let changeData = (row,col,data) => {
	clientData[row][col] = data.detail.data;
	console.log(JSON.stringify(clientData))
	setData()
}

let setData = () => {
	let uuid = inputEl.value;
	fetch(`http://localhost:3000/setlayout/${uuid}`,{
		method:"POST",
		body:JSON.stringify(clientData)
	}).then(x=>{
		console.log(x.status,x.statusText)
	})
}

async function apiCall() {
	let uuid = inputEl.value;
	const res = await fetch(`http://localhost:3000/getlayout/${uuid}`);
	const text = await res.json();

	if (res.ok) {
		clientData = text;
		console.log(clientData)
		return text;
	} else {
		throw new Error(text);
	}
}

const handleClick = () => {
	request = apiCall()
}

onMount(()=>{
	// handleClick()
})

</script>

<style>
.row{
	height: 100px;
}
.search{
	border: 2px solid #008bff;
	padding: 10px;
	font-weight: bold;
	font-size: 20px;
	border-radius: 30px;
	margin-left:10px;
	width: 100%;
}
.getButton{
	padding: 12px 18px;
	border: 2px solid #008bff;
	border-radius: 18px;
	width: min-content;
	height: min-content;
	margin: 10px;
}
.flex-container{
	display: flex;
	align-items: center;
	max-width: 700px;
}
</style>