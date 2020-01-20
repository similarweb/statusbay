const initialState = {};

export function deployment(state = initialState, action) {

  switch (action.type) {
    case 'NEW':
      return action.deployment;    
    default:
      return state
  }
}