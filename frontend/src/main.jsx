import React, { useState, useEffect } from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider, Route, Link, useParams } from 'react-router-dom';
import Grid from '@mui/material/Unstable_Grid2';
import Button from '@mui/material/Button';
import Breadcrumbs from '@mui/material/Breadcrumbs';
import MUILink from '@mui/material/Link';
import Typography from '@mui/material/Typography';
import { useQuery, QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { api } from './api';

const Item = ({ childrens }) => <div>{childrens}</div>;

const Root = () => (
  <>
    <Grid container xs={12}>
      Hello from Gosta!
    </Grid>
    <Grid container xs={12}>
      <Grid xs>
        <Link to={'/runs'}>View runs</Link>
      </Grid>
      <Grid xs>
        <Link to={'/tbd'}>View whatever else(TBD)</Link>
      </Grid>
    </Grid>
  </>
);

const Runs = () => {
  const { data: runs, isLoading } = useQuery(['runs'], () => api.runs.all());

  return (
    <Grid container>
      <Grid xs={12}>
        <nav>
          <Breadcrumbs>
            <MUILink underline="hover" color="inherit" href="/">
              <Link to={'/'}>Home</Link>
            </MUILink>
            <Typography color="text.primary">Runs</Typography>
          </Breadcrumbs>
        </nav>
      </Grid>
      <Grid xs={12}>
        <h1>Runs</h1>
        {isLoading && <h2>Loading...</h2>}
        {runs?.length && (
          <ul>
            {runs.map(({ id, name, status }) => (
              <li>
                <div>
                  <Typography id={id} style={{ display: 'inline', color: status === 'failed' ? 'red' : 'green' }}>
                    {name}
                  </Typography>{' '}
                  <Button variant="outlined">
                    <Link to={`/runs/${id}`}>Open!</Link>
                  </Button>
                </div>
              </li>
            ))}
          </ul>
        )}
      </Grid>
    </Grid>
  );
};

const Suites = () => {
  const { runId } = useParams();
  const { data: run } = useQuery(['run'], () => api.runs.one(runId));
  const { data: suites, isLoading, isError, error } = useQuery(['suites'], () => api.suites.all(runId));

  return (
    <Grid container>
      <Grid xs={12}>
        <nav>
          <Breadcrumbs>
            <MUILink underline="hover" color="inherit" href="/">
              <Link to={'/'}>Home</Link>
            </MUILink>
            <MUILink underline="hover" color="inherit" href="/runs">
              <Link to={'/runs'}>Runs</Link>
            </MUILink>
            {run && <Typography color="text.primary">{run.name}</Typography>}
          </Breadcrumbs>
        </nav>
      </Grid>
      <Grid xs={12}>
        <h1>Suites</h1>
        {isError && <span>Error ocurred: {error.message}</span>}
        {isLoading && <span>Loading....</span>}
        {suites?.length && (
          <ul>
            {suites.map(({ id, name, status }) => (
              <li id={id}>
                <div>
                  <Typography id={id} style={{ display: 'inline', color: status === 'failed' ? 'red' : 'green' }}>
                    {name}
                  </Typography>{' '}
                  <Button variant="outlined">
                    <Link to={`/runs/${runId}/suites/${id}`}>Open!</Link>
                  </Button>
                </div>
              </li>
            ))}
          </ul>
        )}
      </Grid>
    </Grid>
  );
};

const router = createBrowserRouter([
  {
    path: '/',
    element: <Root />,
  },
  {
    path: '/runs',
    element: <Runs />,
  },
  {
    path: '/runs/:runId',
    element: <Suites />,
  },
]);

const queryClient = new QueryClient();

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>
  </React.StrictMode>
);
