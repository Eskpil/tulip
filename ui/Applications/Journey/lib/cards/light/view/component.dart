import 'package:flutter/material.dart';
import 'package:journey/core/models/entity.dart';
import 'package:journey/core/models/state.dart';
import 'package:material_design_icons_flutter/material_design_icons_flutter.dart';

import '../../../core/models/light.dart';

class LightView extends StatelessWidget {
  final Entity entity;
  final EntityState state;

  LightView({ super.key, required this.entity, required this.state });

  @override
  Widget build(BuildContext context) {
    final lightState = LightState.fromParent(state);
    return Container(
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(10),
        color: Colors.black45,
        boxShadow: const [
          BoxShadow(color: Colors.black38, spreadRadius: 3),
        ],
      ),
      padding: const EdgeInsets.all(12),
      child: Container(
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
      )
    );
  }
}