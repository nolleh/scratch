<script lang="ts">
	import fetchChat from '../../../../libs/data';
	// import { fetchChat } from 'libs/data';

	interface User {
		name: string;
	}

	export let friend: User;
	// $: name = friend.name;
	// $: name, reset(friend);

	// function reset(friend: User) {
	// 	message = '';
	//    loadChat(friend.name);
	// }

	// let message = '';

	let chats: string[] = [];
	let loading = false;

	loadChat(friend.name);

	async function loadChat(name: string) {
		loading = true;
		chats = await fetchChat(name);
		loading = false;
	}
</script>

<div class="container">
	<div class="content">
		<div>Talking to {friend.name}</div>

		{#if loading}
			Loading..
		{:else}
			<ul>
				{#each chats as chat}
					<li>{chat}</li>
				{/each}
			</ul>
		{/if}
	</div>
</div>
