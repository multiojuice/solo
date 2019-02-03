import graphene
from flask_graphql import GraphQLView
from flask import Flask

# app initialization
app = Flask(__name__)
app.debug = True

class Query(graphene.ObjectType):
    hello = graphene.String(description='A typical hello world')

    def resolve_hello(self, info):
        return 'World'

schema = graphene.Schema(query=Query)

app.add_url_rule(
    '/graphql',
    view_func=GraphQLView.as_view(
        'graphql',
        schema=schema,
        graphiql=True # for having the GraphiQL interface
    )
)

@app.route('/')
def index():
    return '<p> Hello World</p>'

if __name__ == '__main__':
    app.run()
