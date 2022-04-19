curl -s https://learn.01founders.co/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"cmarcelo963\"}}){id}}"}' | cut -c24-26

