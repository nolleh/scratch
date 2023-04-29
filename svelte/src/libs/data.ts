function fetch(name: string): Promise<string[]> {
	const p = new Promise<string[]>((resolve, _) => {
		setTimeout(() => resolve([`hello, ${name}`, 'how are you doing']), 1000);
	});

	return p;
}

export default fetch;
