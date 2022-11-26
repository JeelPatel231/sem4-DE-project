<label aria-label="seat-label">    
    <input on:click={forward} type="checkbox" bind:this={el} checked={checked || null} prechecked={checked || null}><span />
</label>


<script>
import { createEventDispatcher } from 'svelte';

let el;
export let checked;
const dispatch = createEventDispatcher();

function forward() {
    let value;
    if(el.checked){value = 2}
    else{value = 1}
    console.log("dispatched")
    dispatch('seatDataChange', {data:value});
}
</script>

<style>
label[aria-label="seat-label"]{
    display: inline-block;
}

label[aria-label="seat-label"] > input{
    display: none;
}

label[aria-label="seat-label"] > span::before{
    height: 100px;
    width: 100px;
    content: '';
    background-image: url("/chair.svg");
    background-size: contain;
    display: block;
    position: relative;
    filter: invert(50%) sepia(57%) saturate(2556%) hue-rotate(16deg) brightness(80%) contrast(105%);
}

label[aria-label="seat-label"] > input:checked + span::before{
    filter: invert(50%) sepia(57%) saturate(2556%) hue-rotate(159deg) brightness(91%) contrast(105%);
}

label[aria-label="seat-label"] > input:checked:disabled + span::before{
    filter: invert(25%) sepia(98%) saturate(7483%) hue-rotate(324deg) brightness(104%) contrast(108%);
}

</style>