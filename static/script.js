
const itemInput = document.getElementById('item-input')
const createItemButton = document.getElementById('create-item-button')
const deleteButtons = document.querySelectorAll('.item__delete-button')

async function postItem(event) {
	event.preventDefault()

	const data = {
		title: itemInput.value,
		done: false
	}

	await fetch('http://localhost:3000/api/create', {
		method: "POST",
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify(data)
	})

	itemInput.value = null

	location.reload()
}

async function deleteItem(event) {
	event.preventDefault()

	const key = event.target.getAttribute('data-key')
	const data = { key }

	await fetch('http://localhost:3000/api/delete', {
		method: "DELETE",
		header: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(data)
	})

	location.reload()
}

createItemButton.addEventListener('click', event => postItem(event))

deleteButtons.forEach(button => {
	button.addEventListener('click', event => deleteItem(event))
});
