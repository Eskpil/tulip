import 'package:equatable/equatable.dart';
import 'package:flutter/material.dart';

@immutable
abstract class LightEvent extends Equatable {
  const LightEvent();
}

class LightStarted extends LightEvent {
  @override
  List<Object> get props => [];
}
