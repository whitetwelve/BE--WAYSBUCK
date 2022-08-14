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