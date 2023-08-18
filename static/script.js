
const itemInput = document.getElementById('item-input')
const createItemButton = document.getElementById('create-item-button')
const deleteButtons = document.querySelectorAll('.item__delete-button')

const url = 'https://localhost:3000'

async function preFlight(path, method) {
	await fetch(`${url}${path}`, {
		method: 'OPTIONS',
		headers: {
			"Access-Control-Request-Method": method.toUpperCase(),
			"Access-Control-Request-Headers": "origin, x-requested-with",
			"Origin": `${url}`
		}
	})
}

async function makeRequest(path, method, data) {
	await fetch(`${url}${path}`, {
		method: method.toUpperCase(),
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(data)
	})
}

async function postItem(event) {
	event.preventDefault()

	const path = '/api/create'
	const data = {
		title: itemInput.value,
		done: false
	}

	await preFlight(path, 'POST')
	await makeRequest(path, 'POST', data)

	itemInput.value = null

	location.reload()
}

async function deleteItem(event) {
	event.preventDefault()

	const path = '/api/delete'
	const key = event.target.getAttribute('data-key')
	const data = { key }

	await preFlight(path, 'DELETE')
	await makeRequest(path, 'DELETE', data)

	location.reload()
}

createItemButton.addEventListener('click', event => postItem(event))

deleteButtons.forEach(button => {
	button.addEventListener('click', event => deleteItem(event))
});
