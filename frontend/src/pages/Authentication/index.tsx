import { Container } from "@material-ui/core"
import AppBar from "@material-ui/core/AppBar"
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles"
import Toolbar from "@material-ui/core/Toolbar"
import Typography from "@material-ui/core/Typography"
import React, { useState } from "react"
import Form, { AuthForm } from "./Form"

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    appBar: {
      marginBottom: theme.spacing(2),
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    title: {
      flexGrow: 1,
    },
    formContainer: {
      position: "absolute",
      top: "50%",
      left: "50%",
      transform: "translate(-50%, -50%)",
    },
  })
)

function Authenticate() {
  const classes = useStyles()
  const initialFormState = { login: "", password: "" }
  const authFormState = useState<AuthForm>(initialFormState)

  const onSubmit = (form: AuthForm): void => {
    console.table(form)
  }

  return (
    <>
      <AppBar className={classes.appBar} position="static">
        <Toolbar>
          <Typography variant="h6" className={classes.title}>
            Gauth
          </Typography>
        </Toolbar>
      </AppBar>
      <Container className={classes.formContainer} maxWidth="sm">
        <Form
          onSubmit={onSubmit}
          initialState={initialFormState}
          resetOnSubmit
          formState={authFormState}
        />
      </Container>
    </>
  )
}

export default Authenticate
