import 'package:flutter/material.dart';
import 'package:journey/entity/view/component.dart';
import 'package:journey/entities/bloc/entities_state.dart';

import 'package:flutter_bloc/flutter_bloc.dart';

import '../bloc/entities_bloc.dart';

class EntitiesPage extends StatelessWidget {
  const EntitiesPage({ super.key });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        color: Colors.black87,
        child: const EntityList(),
      ),
    );
  }
}

class EntityList extends StatelessWidget {
  const EntityList({ super.key });

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<EntitiesBloc, EntitiesState>(
      builder: (context, state) {
        if (state is EntitiesLoading) {
          return const CircularProgressIndicator();
        }

        if (state is EntitiesLoaded) {
          List<Widget> cards = [];

          for (final entity in state.entities) {
            cards.add(
              Container(
                margin: const EdgeInsets.all(18),
                child: EntityView(entity: entity),
              ),
            );
          }

          return GridView.count(
            padding: const EdgeInsets.all(12),
            crossAxisCount: 3,
            children: cards,
          );
        }

        return Container();
      }
    );
  }
}