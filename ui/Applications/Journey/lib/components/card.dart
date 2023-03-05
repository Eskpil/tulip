import 'package:flutter/material.dart';
import 'package:material_design_icons_flutter/material_design_icons_flutter.dart';

import '../models/entity.dart';
import '../models/light.dart';

class CardWidget extends StatefulWidget {
  const CardWidget({super.key, required this.entity});

  @override
  _CardWidgetState createState() => _CardWidgetState();

  final Entity entity;
}

class _CardWidgetState extends State<CardWidget> {
  _onBinaryLightPressed() {
    Light.opposite(widget.entity);
  }

  Widget _buildInner() {
    if (widget.entity.kind == "light") {
      return Container(
        margin: const EdgeInsets.only(top: 36),
        child: Column(
          children: [
            IconButton(
              icon: new Icon(MdiIcons.lightbulbCflSpiral),
              color: Colors.white,
              iconSize: 72,
              onPressed: _onBinaryLightPressed,
            ),
            Text(widget.entity.name, style: const TextStyle(color: Colors.grey))
          ],
        ),
      );
    }

    if (widget.entity.kind == "sensor") {
      Icon? icon;

      if (widget.entity.entityMetadata["device_class"] == "temperature") {
        icon = new Icon(MdiIcons.temperatureCelsius);
      }

      if (widget.entity.entityMetadata["device_class"] == "humidity") {
        icon = new Icon(MdiIcons.waterPercent);
      }

      final name = widget.entity.name;
      final unit = widget.entity.entityMetadata["unit_of_measurement"];

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
    return Container(
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(10),
        color: Colors.black45,
        boxShadow: const [
          BoxShadow(color: Colors.black38, spreadRadius: 3),
        ],
      ),
      padding: const EdgeInsets.all(12),
      child: _buildInner(),
    );
  }
}
