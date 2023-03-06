import 'package:flutter/material.dart';
import 'package:equatable/equatable.dart';
import 'package:journey/core/models/light.dart';

@immutable
abstract class LightState extends Equatable {
  const LightState();
}

class LightLoading extends LightState {
  @override
  List<Object> get props => [];
}

class LightLoaded extends LightState {
  final Light light;

  const LightLoaded({ required this.light });

  @override
  List<Object> get props => [light];
}

class LightError extends LightState {
  @override
  List<Object> get props => [];
}
