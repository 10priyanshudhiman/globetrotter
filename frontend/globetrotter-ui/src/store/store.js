import {
    createStore, combineReducers, compose, applyMiddleware,
  } from 'redux';
import {thunk} from 'redux-thunk';
import Reducer from './reducer';
  
  const middlewares = [thunk];
  const composeParams = [applyMiddleware(...middlewares)];
  const store = createStore(
    combineReducers({
      data: Reducer,
    }),
    compose(...composeParams),
  );
  
  export default store;
  