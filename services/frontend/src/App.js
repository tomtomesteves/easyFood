import React, {Component} from 'react';
import "./App.css";
import { ApolloProvider, useMutation } from "@apollo/react-hooks";
import { client } from "./apollo-client";
import { Route } from "react-router";
import { GetAllRestaurants, GetRestaurantsByCategory } from "./gqlComponents/restaurant/RestaurantContainer";
import { useState, useEffect } from "react";


const RestaurantPage = ({match}) => {
  const {
    params: { restaurantId },
  } = match;
  const [isLoading, setIsLoading] = useState(true);
  const [data, setData] = useState();

  useEffect(() => {
  },[restaurantId]);
  return (
    <div>
      <ApolloProvider client={client}>
        <GetRestaurantsByCategory filter={restaurantId} />
      </ApolloProvider>
    </div>
  )
}

function RestaurantByCategory() {
  return (
    <h1>TESTETESTETESTETESTETESTETESTETESTETESTETESTETESTETESTETESTETESTETESTETESTE</h1>
  );
}



function Restaurants() {
  return (
    <div>
      <ApolloProvider client={client}>
        <GetAllRestaurants />
      </ApolloProvider>
    </div>
  );
}


// function App() {
//   var teste = true;
//   if (teste) {
//     return <Restaurants onClick={teste = false} />;
//   }
//   return <restaurantByCategory />;
// }


class App extends React.Component {
  constructor(props) {
    super(props);
  }


  render() {
    let button;
    button = (
      <div>
        <Route exact path="/restaurant" component={Restaurants}>
          <Restaurants />
        </Route>
        <Route path="/restaurant/:restaurantId" component={RestaurantPage} />
      </div>
    );

    return (
      <div>
        {button}
      </div>
    );
  }
}

export default App;
