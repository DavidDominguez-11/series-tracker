import updateRanking from '../utils/updateRanking.js'

const RankingControls = ({ ranking, sortOrder, seriesId, reRenderTable }) => {
  const container = document.createElement('div')
  container.classList.add('ranking-controls')
  const num = document.createElement('span')
  num.textContent = ranking
  container.appendChild(num)

  // Upvote button
  const upButton = document.createElement('button')
  upButton.textContent = 'Up'
  upButton.addEventListener('click', async () => {
    try {
      const direction = sortOrder === 'asc' ? 'downvote' : 'upvote'
      await updateRanking(seriesId, direction)
      console.log('Ranking updated')
      reRenderTable() // Ensure this is a valid function
    } catch (error) {
      console.error('Failed to update ranking:', error)
    }
  })
  container.appendChild(upButton)

  // Downvote button
  const downButton = document.createElement('button')
  downButton.textContent = 'Down'
  downButton.addEventListener('click', async () => {
    try {
      const direction = sortOrder === 'asc' ? 'upvote' : 'downvote'
      await updateRanking(seriesId, direction)
      console.log('Ranking updated')
      reRenderTable() // Ensure this is a valid function
    } catch (error) {
      console.error('Failed to update ranking:', error)
    }
  })
  container.appendChild(downButton)

  return container
}

export default RankingControls

