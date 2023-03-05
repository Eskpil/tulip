import 'package:flutter/material.dart';
import 'package:equatable/equatable.dart';

@immutable
abstract class EntitiesEvent extends Equatable {
  const EntitiesEvent();
}

class EntitiesStarted extends EntitiesEvent {
  @override
  List<Object> get props => [];
}