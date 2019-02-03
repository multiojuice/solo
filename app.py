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

class ChangeFile(graphene.Mutation):
    class Arguments:
        name = graphene.String(required=True)
        text = graphene.String(required=True)

    ok=graphene.Boolean()

    def mutate(self, info, name, text):
        with open(name, 'w') as opened_file:
            opened_file.write(text)
        ok = True
        return ChangeFile(ok=ok)

class Mutation(graphene.ObjectType):
    change_file = ChangeFile.Field()

schema = graphene.Schema(query=Query, mutation=Mutation)

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
