import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import DeleteIcon from '@material-ui/icons/Delete';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../services/apis';
import { EntUser } from '../../services/models/EntUser'; // import interface User
//import { EntClub } from '../../services/models/EntClub'; // import interface Club
//import { EntClubEdgesFromJSON, EntClubEdgesFromJSONTyped, EntUserEdges, mapValues } from '../../services';

const useStyles = makeStyles({
  table: {
    minWidth: 650,
  },
});

export default function ComponentsTable() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [users, setUsers] = useState<EntUser[]>([]);

  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getUsers = async () => {
      const res = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(res);
      console.log(res);
    };
    getUsers();;
  }, [loading]);

  const deleteUsers = async (id: number) => {
    const res = await api.deleteUser({ id: id });
    setLoading(true);
  };

  const deleteInformations = async (id: number) => {
    const res = await api.deleteInformation({ id: id });
    setLoading(true);
  };

  return (
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell align="center">ID</TableCell>
            <TableCell align="center">Email</TableCell>
            <TableCell align="center">Club</TableCell>
            <TableCell align="center">Role</TableCell>
            <TableCell align="center">Information Name</TableCell>
            <TableCell align="center">Manage</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {users === undefined
            ? null
            : users.map((item) => (
            <TableRow key={item.id}>
              <TableCell align="center">{item.id}</TableCell>
              <TableCell align="center">{item.email}</TableCell>
              <TableCell align="center">{item.edges?.club?.cLUBNAME}</TableCell>
              <TableCell align="center">{item.edges?.role?.tITLE}</TableCell>
              <TableCell align="center">{item.edges?.information?.nAME}</TableCell>
              <TableCell align="center">
                <Button
                  onClick={() => {
                    if(item.id === undefined){
                      return;
                    }
                    deleteUsers(item.id);
                    if(item.edges?.information?.id === undefined){
                      return;
                    }
                    deleteInformations(item.edges?.information?.id);
                  }}
                  style={{ marginLeft: 10 }}
                  variant="contained"
                  color="secondary"
                  startIcon={<DeleteIcon />}
                >
                  Delete
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
