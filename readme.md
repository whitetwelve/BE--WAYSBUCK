<!-- RUN NODEMON -->
nodemon --exec go run main.go

<!-- FLOW  -->
repo --> models --> handler --> routes --> main

<!-- RELATION TABLES -->
USER --> PROFILE {HASONE}
PROFILE --> USER {BELONGSTO}
USER && PROFILE === ONE TO ONE 

USER --> PRODUCT {HAS MANY}
PRODUCT --> USER {BELONGSTOMANY}

PRODUCT --> TOPING {HAS MANY}
TOPING --> PRODUCT {BELONGSTO MANY}

<!-- TRANSACTION -->
SNAP : POP UP PEMBAYARAN PADA MIDTRANS
PASSWORD MIDTRANS : Fuadazkia17

<!-- BE DEPLOY -->
https://waysbuck-fuad.herokuapp.com/api/v1/

<!-- FE DEPLOY -->
https://waysbuck-dumbways.netlify.app/