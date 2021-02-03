import React, {Component} from 'react';
import "./App.css";
import { ApolloProvider, useMutation } from "@apollo/react-hooks";
import { client } from "./apollo-client";
import { Route } from "react-router";
import { Link } from "react-router-dom";
import {
    GetAllRestaurants,
    GetRestaurantsByID
} from "./gqlComponents/restaurant/RestaurantContainer";
import {
    GetAllCategories,
    GetCategoryByID
} from "./gqlComponents/category/CategoryContainer";
import { useState, useEffect } from "react";
import NavBar from './NavBar'


const CategoryPage = ({match}) => {
    const {
        params: { categoryId },
    } = match;
    const [isLoading, setIsLoading] = useState(true);
    const [data, setData] = useState();

    useEffect(() => {
    },[categoryId]);
    return (
        <div>
            <ApolloProvider client={client}>
                <GetCategoryByID filter={categoryId} />
            </ApolloProvider>
        </div>
    )
}

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
        <GetRestaurantsByID filter={restaurantId} />
      </ApolloProvider>
    </div>
  )
}

function Categories() {
    return (
        <div>
            <ApolloProvider client={client}>
                <GetAllCategories />
            </ApolloProvider>
        </div>
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

function HomePage() {
  return (
    <div>
      <ApolloProvider client={client}>
        <Link to="/restaurant">Restaurantes</Link>
        <Link to="/category">Categorias</Link>
        <Link to="/dish">Pratos</Link>
      </ApolloProvider>
    </div>
  );
}

class App extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    let button;
    button = (
      <div>
        <div>
          <NavBar />
        </div>
        <Route exact path="/restaurant" component={Restaurants}>
          <Restaurants />
        </Route>
        <Route path="/restaurant/:restaurantId" component={RestaurantPage} />

        <Route exact path="/category" component={Categories}>
          <Categories />
        </Route>
        <Route path="/category/:categoryId" component={CategoryPage} />

        <Route exact path="/" component={HomePage}>
          <HomePage />
        </Route>

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