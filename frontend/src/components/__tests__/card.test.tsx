import { render, screen } from '@testing-library/react'
import '@testing-library/jest-dom'
import CardComponent from '../card'

describe('CardComponent', () => {
  const cardData = {
    id: 1,
    name: 'John Doe',
    email: 'john.doe@example.com',
  }

  test('renders card with correct details', () => {
    render(<CardComponent card={cardData} />)

    // Check if the component renders the card details correctly
    expect(screen.getByText(`Id: ${cardData.id}`)).toBeInTheDocument()
    expect(screen.getByText(cardData.name)).toBeInTheDocument()
    expect(screen.getByText(cardData.email)).toBeInTheDocument()
  })

  test('has correct styles', () => {
    const { container } = render(<CardComponent card={cardData} />)

    // Check if the component has the correct styles
    const cardElement = container.firstChild as HTMLElement

    expect(cardElement).toHaveClass('bg-white')
    expect(cardElement).toHaveClass('shadow-lg')
    expect(cardElement).toHaveClass('rounded-lg')
    expect(cardElement).toHaveClass('p-2')
    expect(cardElement).toHaveClass('mb-2')
    expect(cardElement).toHaveClass('hover:bg-gray-100')
  })
})
