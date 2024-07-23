import { render, screen, fireEvent, waitFor, act } from '@testing-library/react'
import '@testing-library/jest-dom'
import axios from 'axios'
import MockAdapter from 'axios-mock-adapter'
import UserInterface from '../user-interface'

// Setup axios mock
const mock = new MockAdapter(axios)

describe('UserInterface', () => {
  const apiUrl = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8000'

  afterEach(() => {
    mock.reset()
  })

  test('fetches and displays users', async () => {
    // Mock the GET request
    mock.onGet(`${apiUrl}/api/go/users`).reply(200, [
      { id: 1, name: 'John Doe', email: 'john@example.com' },
      { id: 2, name: 'Jane Doe', email: 'jane@example.com' },
    ])

    render(<UserInterface backendName='go' />)

    // Wait for the data to be fetched and rendered
    await waitFor(() => {
      expect(screen.getByText('John Doe')).toBeInTheDocument()
      expect(screen.getByText('Jane Doe')).toBeInTheDocument()
    })
  })

  test('creates a new user', async () => {
    // Mock GET and POST requests
    mock.onGet(`${apiUrl}/api/go/users`).reply(200, [])
    mock
      .onPost(`${apiUrl}/api/go/users`)
      .reply(201, { id: 1, name: 'New User', email: 'new@example.com' })

    render(<UserInterface backendName='go' />)

    fireEvent.change(screen.getByPlaceholderText('Name'), {
      target: { value: 'New User' },
    })
    fireEvent.change(screen.getByPlaceholderText('Email'), {
      target: { value: 'new@example.com' },
    })
    fireEvent.click(screen.getByText('Add User'))

    await waitFor(() => {
      expect(screen.getByText('New User')).toBeInTheDocument()
    })
  })

  test('updates a user', async () => {
    // Mock GET request to fetch users
    mock
      .onGet(`${apiUrl}/api/go/users`)
      .reply(200, [{ id: 1, name: 'Old User', email: 'old@example.com' }])

    // Mock PUT request to update the user
    mock.onPut(`${apiUrl}/api/go/users/1`).reply(200)

    // Mock GET request to fetch users after update
    mock
      .onGet(`${apiUrl}/api/go/users`)
      .replyOnce(200, [
        { id: 1, name: 'Updated User', email: 'updated@example.com' },
      ])

    render(<UserInterface backendName='go' />)

    // Wait for the initial user to be rendered
    await waitFor(() => {
      expect(screen.getByText('Old User')).toBeInTheDocument()
    })

    // Fill out and submit the update form
    fireEvent.change(screen.getByPlaceholderText('User Id'), {
      target: { value: '1' },
    })
    fireEvent.change(screen.getByPlaceholderText('New Name'), {
      target: { value: 'Updated User' },
    })
    fireEvent.change(screen.getByPlaceholderText('New Email'), {
      target: { value: 'updated@example.com' },
    })

    const updateButton = screen.getByText('Update User')
    fireEvent.click(updateButton)

    // Wait for the user to be updated
    await waitFor(() => {
      expect(screen.getByText('Updated User')).toBeInTheDocument()
      expect(screen.queryByText('Old User')).not.toBeInTheDocument()
    })
  })

  test('deletes a user', async () => {
    // Mock GET request to fetch users
    mock
      .onGet(`${apiUrl}/api/go/users`)
      .reply(200, [
        { id: 1, name: 'User to Delete', email: 'delete@example.com' },
      ])

    // Mock DELETE request
    mock.onDelete(`${apiUrl}/api/go/users/1`).reply(200)

    // Mock GET request to verify user deletion
    mock.onGet(`${apiUrl}/api/go/users`).replyOnce(200, []) // Ensure it returns empty list after deletion

    render(<UserInterface backendName='go' />)

    // Wait for the user to be rendered
    await waitFor(() => {
      expect(screen.getByText('User to Delete')).toBeInTheDocument()
    })

    // Click delete button
    const deleteButton = screen.getByText('Delete User')
    fireEvent.click(deleteButton)

    // Wait for the list to be updated
    await waitFor(() => {
      expect(screen.queryByText('User to Delete')).not.toBeInTheDocument()
    })
  })
})
