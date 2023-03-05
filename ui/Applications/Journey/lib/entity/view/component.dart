import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:journey/entity/bloc/entity_state_event.dart';
import 'package:material_design_icons_flutter/material_design_icons_flutter.dart';

import 'package:journey/entities/models/entity.dart';
import 'package:journey/entity/bloc/entity_state_bloc.dart';
import 'package:journey/entity/bloc/entity_state_state.dart';
import 'package:journey/entity/models/entity_state.dart';

import '../../core/models/light.dart';
import '../../repositories/entity.dart';

class EntityView extends StatelessWidget {
  final Entity entity;

  const EntityView({ super.key, required this.entity });

  Widget _buildInner(EntityState state) {
    if (entity.kind == "light") {
      final lightState = LightState.fromParent(state);

      return Container(
        margin: const EdgeInsets.only(top: 36),
        child: Column(
          children: [
            IconButton(
              icon: new Icon(MdiIcons.lightbulbCflSpiral),
              color: lightState.state == "ON" ? Colors.white : Colors.grey,
              iconSize: 72,
              onPressed: () {
                Light.opposite(entity);
              },
            ),
            Text(entity.name, style: const TextStyle(color: Colors.grey))
          ],
        ),
      );
    }

    if (entity.kind == "sensor") {
      Icon? icon;

      if (entity.entityMetadata["device_class"] == "temperature") {
        icon = new Icon(MdiIcons.temperatureCelsius);
      }

      if (entity.entityMetadata["device_class"] == "humidity") {
        icon = new Icon(MdiIcons.waterPercent);
      }

      final name = entity.name;
      final unit = entity.entityMetadata["unit_of_measurement"];

      return Container(
        margin: const EdgeInsets.only(top: 36),
        child: Column(
          children: [
            Icon(
              icon!.icon,
              color: Colors.white,
              size: 72,
            ),
            Text("$name 22$unit", style: const TextStyle(color: Colors.grey)),
          ],
        ),
      );
    }

    return Container();
  }

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (_) => EntityStateBloc(entityRepository: EntityRepository(), entity: entity)..add(EntityStateStarted()),
      child: BlocBuilder<EntityStateBloc, EntityStateState>(
        builder: (context, state) {
          if (state is EntityStateLoading) {
            return const CircularProgressIndicator();
          }

          if (state is EntityStateUpdated) {
            return Container(
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(10),
                color: Colors.black45,
                boxShadow: const [
                  BoxShadow(color: Colors.black38, spreadRadius: 3),
                ],
              ),
              padding: const EdgeInsets.all(12),
              child: _buildInner(state.state),
            );
          }

          return const Text("Something might have gone wrong");
        },
      ),
    );


  }
}